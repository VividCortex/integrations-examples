# Verify webhook signature 

Example code snippets to verify Solarwinds DPM webhook integrations.
More info available at https://docs.vividcortex.com/how-to-use-vividcortex/integrations/#generic-webhook


### Go
```
go build
./verify-webhook-signature
```

### Python
```
python3 main.py
```

### Node
```
node main.js
```

### Example curl to simulate DPM alerts to your webhook
```
curl --location --request POST 'localhost:1337/hook' \
--header 'X-VividCortex-Signature: 19f3cbef11b0f9a4bb2221178436664365346f04' \
--header 'Content-Type: application/json' \
--data-raw '{
    "environment_name": "Default",
    "event_metric": "os.cpu.total_us",
    "event_type": "Threshold Alert",
    "event_url": "https://app.vividcortex.com/Default/events?from=1526303681&filterFrom=1526303856&filterUntil=1526303981&hosts=id=1",
    "event_uuid": "abadabadabadabadabadabadabadabad",
    "event_level": "warn",
    "host_arch": "amd64",
    "host_description": "Server lubuntu-1604 version 4.10.0-42-generic",
    "host_name": "lubuntu-1604",
    "host_os": "linux",
    "host_type": "os",
    "org_name": "VividCortex",
    "org_nick": "vividcortex",
    "event_message": "Metric os.cpu.total_us at 280.0K, greater than trigger threshold of 10.0 for 1 second",
    "timestamp": 1526303881,
    "alert_id": 2,
    "alert_name": "Threshold Alert",
    "alert_url": "https://app.vividcortex.com/Default/settings/alerts?aid=2",
    "integration_id": 2,
    "integration_name": "Test Webhook",
    "integration_url": "https://app.vividcortex.com/Default/settings/integrations?duid=2"
}'
```
