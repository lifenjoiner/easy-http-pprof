When you need http pprof:
1. Put me together with your main.
2. Build with `pprof` added to `-tags`.
3. Add `-X main.PPROF_SEVER=` with a new address to `-ldflags` to change the server address.
