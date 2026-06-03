# Как вызвать API из PowerShell

Ниже приведён простой пример, как отправить `POST`‑запрос с телом `{"message":"hello"}` к вашему серверу на Go.

```powershell
# Создаём JSON‑строку. PowerShell 5+ имеет cmdlet ConvertTo-Json.
$body = @{ message = "hello" } | ConvertTo-Json

# Отправляем запрос
$response = Invoke-RestMethod \
    -Uri "http://localhost:8080/api" \
    -Method Post \
    -ContentType "application/json" \
    -Body $body

# Выводим полученный JSON
$response | Format-List
```

Если вы используете старую версию PowerShell (до 3.0), можно воспользоваться `curl`‑альтернативой `Invoke-WebRequest`:

```powershell
$body = "{\"message\": \"hello\"}"
$response = Invoke-WebRequest \
    -Uri "http://localhost:8080/api" \
    -Method Post \
    -Headers @{ "Content-Type" = "application/json" } \
    -Body $body

# В `Invoke-WebRequest` тело ответа находится в свойствах `Content` и `RawContent`.
Write-Output $response.Content
```

## Почему вы получаете `invalid2 request`?

- **Пустое тело** – `Invoke-RestMethod` без `-Body` по‑умолчанию не посылает тело.
- **Неверный JSON** – убедитесь, что двойные кавычки экранированы, как показано выше.
- **Отсутствует заголовок `Content-Type: application/json`** – без него сервер может считать запрос неполным.

С помощью приведённого кода запрос будет корректным и сервер вернёт JSON‑ответ:

```json
{"Message":"Response: hello"}
```
