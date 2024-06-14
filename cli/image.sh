pwd=$(pwd)

bin/toast64.exe \
  --app-id "Example App" \
  --title "Hello World" \
  --message "Lorem ipsum dolor sit amet, consectetur adipiscing elit." \
  --icon "$(pwd)/sonic.jpg" \
  --audio "default" --loop \
  --duration "long" \
  --activation-arg "https://google.com" \
  --image "${pwd}/banner.jpg"
