# ethsign
Sign ethereum transactions

Notes
 - `--amount` is in `finney` (ie. `milliether`).
 - `--price` is in `gwei` (ie. `nanoether`).
 - addresses are pure hex (no `0x`)
 - BYON (bring your own nonce)

```
ethsign sign --keydir ~/.ethereum/keystore --from 9e0b9ddba97dd4f7addab0b5f67036eebe687606 --to 37a9679c41e99db270bda88de8ff50c0cd23f326  --amount 10 --price 10 --nonce 20
```
