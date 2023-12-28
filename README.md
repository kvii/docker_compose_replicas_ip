# IP 问题复现

1. 执行 `docker compose up -d`，启动工程。
2. 执行 `curl http://localhost`，观察响应的 name 和 hostname。
3. 编辑 compose.yaml，将 app1 和 app2 的 tag 改成 0.0.2，replica 改成 5。
4. 执行 `docker compose up -d` 更新工程。
5. 再次执行 `curl http://localhost`，响应的 name 属性变为 app2。说明请求被 nginx 转发到了错误的容器上。