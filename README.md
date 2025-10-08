# Docker Quick Guide - Lyvora Server

## Building the Image

### Linux
```bash
sudo docker build -t lyvora-server .
```

### Windows
```powershell
docker build -t lyvora-server .
```

## Running the Service

### Linux
```bash
# Start the service
sudo docker-compose up -d

# Or rebuild and start
sudo docker-compose up --build -d
```

### Windows (PowerShell/CMD)
```powershell
# Start the service
docker-compose up -d

# Or rebuild and start
docker-compose up --build -d
```

## Stopping the Service

### Linux
```bash
sudo docker-compose down
```

### Windows
```powershell
docker-compose down
```

## Checking Status

### Linux
```bash
# View running containers
sudo docker ps

# View logs
sudo docker-compose logs -f
```

### Windows
```powershell
# View running containers
docker ps

# View logs
docker-compose logs -f
```

## Accessing the Service

Once running, access your API at: **http://localhost:8080**

## Testing the Service

### Linux
```bash
curl http://localhost:8080/api/tracks
```

### Windows (PowerShell)
```powershell
curl http://localhost:8080/api/tracks
```

### Windows (CMD)
```cmd
curl http://localhost:8080/api/tracks
```
