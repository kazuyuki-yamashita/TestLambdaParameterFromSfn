# TestLambdaParameterFromSfn

StepFunctions から　 Lambda 　への変数を渡す動作確認

1.  Lambda を作成する

2.  Lambda を呼び出す下記の　 StepFunctions を作成する

    ![StepFunctions](./src/img/img1.png)

    StepFunctions の定義では、Lambda に渡す変数を下記のように、Parameters の Payload に渡す

    ```json
     "Parameters": {
        "Payload": {
          "InvokeNumber": 1
        }
    ```

3.  Lambda 関数を作成する(handler.go)

    - 受け取る変数名を構造体として定義する

    ```go
    type Event struct {
        InvokeNumber string `json:"InvokeNumber"`
    }
    ```

    - 受け取った変数を呼び出す

    ```go
    func Handler(ctx context.Context, event Event) {
        fmt.Println("Hello Sfn!!")
        fmt.Println("InvokeNumber: ", event.InvokeNumber)
    }
    ```

4.  ビルド・デプロイする

    ```sh
    make
    ```

# 動作確認

StepFunctions を実行し、Lambda のログから、変数`InvokeNumber`に代入した値が出力されている

![result](./src/img/img2.png)

# 参考資料

https://qiita.com/yukpiz/items/269277a97053237a6980
