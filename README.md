**Quick Guide to using the tool**
Open your terminal or command prompt and follow these steps:

1. **Pull Code:**
   ```bash
   git clone -b feat/miner-base https://github.com/eternalai-org/eternal.git
   ```
   
2. **Update Environment:**
   Edit `env/config.env` with your settings.

3. **Start Service:**

    ```bash
    docker-compose up -d
    ```
    
4. **Check docker container status:**
    ```bash
    docker ps
    ```
    
5. **View docker container logs:**
    ```bash
    docker logs -f <container-name>
    ```


