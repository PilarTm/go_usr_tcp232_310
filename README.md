**Supported device**
* USR-TCP232-T
* USR-TCP232-2
* USR-TCP232-S
* USR-TCP232-D
* USR-TCP232-24
* USR-TCP232-200
* USR-TCP232-300
* USR-TCP232-204


**The command length is 8 bits, detail as follow table. The demo bytes are in hex mode:**

| Name                       | Packet header | Band rate                                | UART bits setting                             | Check sum                 |
|----------------------------|---------------|------------------------------------------|-----------------------------------------------|---------------------------|
| Bytes                      | 3             | 3                                        | 1                                             | 1                         |
| Description                | Three bytes   | Band rate in hex mode, High byte  first. | Parity/data/stop settings, see follow  table. | Check sum of last 4 bytes |
| For example (115200,N,8,1) | 55 AA 55      | 01 C2 00                                 | 83                                            | 83                        |
| For example (9600,N,8,1)   | 55 AA 55      | 00 25 80                                 | 83                                            | 83                        |

```
Test bits
55AA5501C2008346 For 115200 N,8,1
55AA550025808328 For 9600 N,8,1
```


**Appendix: UART bits setting detail.**

| Bit | Description   | Value | Description       |
|-----|---------------|-------|-------------------|
| 1:0 | Data bits     | 00    | 5 bits            |
|     |               | 01    | 6 bits            |
|     |               | 10    | 7 bits            |
|     |               | 11    | 8 bits            |
| 2   | Stop bits     | 0     | 1 bits            |
|     |               | 1     | 2 bits            |
| 3   | Parity enable | 0     | Not enable Parity |
|     |               | 1     | Enable Parity     |
| 5:4 | Parity type   | 00    | ODD               |
|     |               | 01    | EVEN              |
|     |               | 10    | Mark              |
|     |               | 11    | Clear             |
| 8:6 | Not used      | 000   | Please fill 0     |