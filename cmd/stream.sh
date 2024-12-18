#!/bin/bash

# Define the segment directory and playlist file
SEGMENT_DIR="./segments"
PLAYLIST_FILE="$SEGMENT_DIR/playlist.m3u8"
UPLOADED_PLAYLIST=false
UPLOADED_SEGMENTS_FILE="$SEGMENT_DIR/uploaded_segments.txt"

# Ensure the temporary file exists
touch "$UPLOADED_SEGMENTS_FILE"

# Function to start the live stream with FFmpeg
start_live_stream() {
    ffmpeg -f avfoundation -framerate 30 -i "0" -t 60 -vsync vfr -c:v libx264 -preset veryfast -c:a aac -b:a 128k \
    -f hls -hls_time 10 -hls_playlist_type event \
    -hls_segment_filename "$SEGMENT_DIR/segment_%03d.ts" \
    -hls_playlist_type event \
    -hls_flags independent_segments \
    "$PLAYLIST_FILE" &
}

# Function to check if a segment has already been uploaded
is_uploaded() {
    local segment="$1"
    grep -qx "$segment" "$UPLOADED_SEGMENTS_FILE"
}

# Function to process and upload segments
process_segment() {
    echo "Processing segments"
    for segment_file in "$SEGMENT_DIR"/*.ts; do
        if [[ -f "$segment_file" ]]; then
            local segment_name=$(basename "$segment_file")
            echo "Processing segment: $segment_name"

            if is_uploaded "$segment_name"; then
                echo "Already uploaded: $segment_name"
            else
                echo "New segment detected: $segment_name"
                echo "Uploading $segment_name to S3..."
                # Uncomment the line below to actually upload
                aws s3 cp "$segment_file" "s3://${bucket_name}/segments/$segment_name"
                echo "$segment_name" >> "$UPLOADED_SEGMENTS_FILE"
            fi
        fi
    done

    echo "Uploading playlist.m3u8 to S3..."
    # Uncomment the line below to actually upload
    aws s3 cp "$PLAYLIST_FILE" "s3://${bucket_name}/segments/playlist.m3u8"
    UPLOADED_PLAYLIST=true
}

# Start the live stream
start_live_stream
echo "FFmpeg started with PID: $!"

# Loop every 5 seconds to process segments while FFmpeg is running
while true; do
    FFMPEG_PID=$(pgrep ffmpeg)

    if [[ -z "$FFMPEG_PID" ]]; then
        echo "FFmpeg process not found. Exiting loop."
        break
    else
        echo "Checking for new segments..."
        process_segment
    fi

    sleep 5
done


process_segment
echo "FFmpeg process has ended."

