| Name                       | Packet header | Band rate                                | UART bits setting                             | Check sum                 |
|----------------------------|---------------|------------------------------------------|-----------------------------------------------|---------------------------|
| Bytes                      | 3             | 3                                        | 1                                             | 1                         |
| Description                | Three bytes   | Band rate in hex mode, High byte  first. | Parity/data/stop settings, see follow  table. | Check sum of last 4 bytes |
| For example (115200,N,8,1) | 55 AA 55      | 01 C2 00                                 | 83                                            | 83                        |
| For example (9600,N,8,1)   | 55 AA 55      | 00 25 80                                 | 83                                            | 83                        |


