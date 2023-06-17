package gin

import (
	"aramen-api/srv/logs"
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Gin struct {
	*gin.Engine
	shutdown   chan bool
	server     *http.Server
	serverPort int
	running    bool
	lock       sync.Mutex
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}

func (g *Gin) Port() int {
	g.lock.Lock()
	defer g.lock.Unlock()

	return g.serverPort
}

func (g *Gin) IsRunning() bool {
	g.lock.Lock()
	defer g.lock.Unlock()

	return g.running
}

func (g *Gin) Run(addr ...string) {
	g.server = newHTTPServerAddr(addr[0])
	g.server.Handler = h2c.NewHandler(g.Engine, &http2.Server{
		MaxConcurrentStreams: 1000,
	})
	g.shutdown = make(chan bool)

	err := g.server.Serve(g.listen())
	if errors.Is(err, http.ErrServerClosed) {
		<-g.shutdown
		logs.Info().Msg("HTTP Server shutdown")
	} else {
		logs.Error().Err(err).Msg("HTTP Server cannot be started")
	}
}

func newHTTPServerAddr(addr string) *http.Server {
	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 30 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      300 * time.Second,
	}
	return srv
}

func (g *Gin) listen() net.Listener {
	g.lock.Lock()
	defer g.lock.Unlock()

	logs.Info().Value("listen_addr", g.server.Addr).Msg("HTTP Server is starting and listening")
	ln, err := net.Listen("tcp", g.server.Addr)
	if err != nil {
		panic(err)
	}
	g.running = true
	g.serverPort = ln.Addr().(*net.TCPAddr).Port

	return ln
}

func (g *Gin) ShutdownGracefully() {
	g.lock.Lock()
	defer g.lock.Unlock()
	if !g.running {
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	if err := g.server.Shutdown(ctx); err != nil {
		logs.Error().Err(err).Msg("HTTP Server shutdown with error")
	}

	close(g.shutdown)
	g.running = false
}

func NewServer() *Gin {
	engine := gin.New()
	srv := &Gin{Engine: engine}
	srv.Use(gin.Logger())
	srv.Use(gin.Recovery())
	srv.Use(cors.Default())
	return srv
}
