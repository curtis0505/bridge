# health 체크
[GET] /health
## Response
- 성공
    ```json
    {
        "result": 0
    }
    ```

# validator 주소(체인별)
[GET] /:chainSymbol/validator/address
## Response
- 성공
    ```json
    {
      "result": 0,
      "address": "0xA54042710408c5D87eC8B898886b70F1e299534F"
    }
    ```

# validator 주소(모든 체인)
[GET] /validator
## Response
- 성공
    ```json
    {
        "result": 0,
        "validator": [
            {
                "chainSymbol": "KLAY",
                "address": "0xA54042710408c5D87eC8B898886b70F1e299534F"
            },
            {
                "chainSymbol": "MATIC",
                "address": "0xA54042710408c5D87eC8B898886b70F1e299534F"
            }
        ]
    }
    ```

# 트랜잭션 복구(재시도)
[POST] /:chainSymbol/recover
## Request
- txHash: Deposit, Burn 이벤트가 발생한 트랜잭션 해시
    ```json
    {
        "txHash": "0x6b4afa469788a1e3788aac9008a4009718a9effd92cf6576875d4d78a1734c3b"
    }
    ```

## Response
- 성공
    ```json
    {
        "result": 0
    }
    ```
- 트랜잭션이 조회되지 않음
    ```json
    {
      "result": 10000
    }
    ```
- 트랜잭션이 펜딩중임
    ```json
    {
      "result": 10001
    }
    ```
- 이미 처리된 트랜잭션 일때
    ```json
    {
      "result": 10002
    }
    ```

# 트랜잭션 스피드업
[POST] /:chainSymbol/speedup
## Request
- txHash: 해당 validator 가 실행한 submitTransaction 해시
    ```json
    {
        "txHash": "0x6b4afa469788a1e3788aac9008a4009718a9effd92cf6576875d4d78a1734c3b"
    }
    ```

## Response
- 성공
    ```json
    {
        "result": 0
    }
    ```
- 트랜잭션이 조회되지 않음
    ```json
    {
      "result": 10000
    }
    ```
- 이미 처리된 트랜잭션 일때
    ```json
    {
      "result": 10002
    }
    ```

# 트랜잭션 취소
[DELETE] /:chainSymbol/transaction
## Request
- nonce: 취소할 트랜잭션의 nonce
- gasPrice: gasPrice
- gasTip: gasTip, 옵션값 클레이튼에서는 입력 안해도됨
    ```json
    {
        "nonce": 366,
        "gasPrice": 10000000000,
        "gasTip": 10000000000
    }
    ```

## Response
- 성공
    ```json
    {
        "result": 0
    }
    ```
- 취소 실패
    ```json
    {
      "result": 10000
    }
    ```
- 이미 처리된 트랜잭션 일때
    ```json
    {
      "result": 10002
    }
    ```
