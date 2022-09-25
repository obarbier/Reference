## Definition
one can capture [UNIX signals](https://www.tutorialspoint.com/unix/unix-signals-traps.htm) using goland signal.Notify

```
// Handle sigterm and await termChan signal
termChan := make(chan os.Signal)
signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

go func() {
	<-termChan // Blocks here until interrupted
	log.Print("SIGTERM received. Shutdown process initiated\n")
	httpServer.Shutdown(context.Background())
}()

```

finishing running job can be done usinge async waitgroup and aor queue system

## Glossary

## Reference
* [Rodrigo - Blog](https://www.rodrigoaraujo.me/posts/golang-pattern-graceful-shutdown-of-concurrent-events/)