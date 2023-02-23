# NSQ

* First, run these commands to start the consumer

    ```bash
  go build
  ./consumer --topic "topic_name" --lookupd "127.0.0.1:4161"
    ```