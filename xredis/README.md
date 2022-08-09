# Redis Client 
A Redis Client support single and cluster mode. 

## Usage 

## Redis new client with config 
```go 
config := xredis.Config{
    SingleMode: true, //Single or cluster 
    URL:        "localhost:6379",
}
client, _, err := xredis.Init(config)
if err != nil {
    log.WithError(err).Fatalf("GetRedisClient returns error: %s", err.Error())
}
```

## Redis new single client 
```go 
client, err := GetRedisClient(config.URL, config.Password)
client := NewClient(redisClient)
```

## Redis new cluster client 
```go 
client, err := NewClusterClient(config.URL, config.Password)
client := NewClient(redisClient)
```