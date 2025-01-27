#!/bin/bash

DIRECTORY_TO_MONITOR="/home/hijokaidan/PC/jardin-pc/internal/views/"
SCRIPT_1="./dev-run.sh"

kill_process_by_port() {
    local port=42069
    local pid=$(lsof -t -i:$port)
    echo $pid
    echo "Killing process on port $port with PID: $pid"
    kill -9 $pid
}

handle_file_change() {
    echo "File change detected in $DIRECTORY_TO_MONITOR. Stopping running process and re-executing $SCRIPT_1"

    cd ..
    ~/go/bin/templ generate
    npx tailwindcss -i tailwind.base.css -o resources/static/css/tailwind.css

    kill_process_by_port
    cd internal
    go run main.go &
}

monitor_directory() {
    local last_modified=$(find "$DIRECTORY_TO_MONITOR" -type f ! -name "*.go" -exec stat -c '%Y %n' {} \; | sort -nr | head -n 1 | cut -d ' ' -f 1)
    handle_file_change
    while true; do
        local current_modified=$(find "$DIRECTORY_TO_MONITOR" -type f  ! -name "*.go" -exec stat -c '%Y %n' {} \; | sort -nr | head -n 1 | cut -d ' ' -f 1)

        if [ "$current_modified" -gt "$last_modified" ]; then
            handle_file_change
            last_modified=$current_modified
        fi

        sleep 1
    done
}

monitor_directory

