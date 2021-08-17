# Coingecko price Alert

## Build

```
go build -o coing .
```

## Example

If price is higher than
```bash
./coing -coin=btc -up=60000
```

If price is less than
```bash
./coing -coin=btc -down=30000
```

**Notify by email**

```bash
./p2pc -coin=btc -up=50000 \ 
    -mail.host=mail.example.com \
    -mail.port=587 \
    -mail.user=user@example.com \
    -mail.pass=xxxx \
    -mail.from=noreply@example.com \
    -mail.to=me@gmail.com
```


