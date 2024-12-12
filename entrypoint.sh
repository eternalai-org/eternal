#!/bin/bash

# Start Ollama in the background.
/bin/ollama serve &
# Record Process ID.
pid=$!

# Pause for Ollama to start.
sleep 5

echo "🔴 Retrieve hf.co/lmstudio-community/INTELLECT-1-Instruct-GGUF:Q8_0 model..."
ollama run hf.co/lmstudio-community/INTELLECT-1-Instruct-GGUF:Q8_0
echo "🟢 Done!"

# Wait for Ollama process to finish.
wait $pid