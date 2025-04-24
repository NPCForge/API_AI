URL="http://api:3000/health"
EXPECTED="OK"
TIMEOUT=30

echo "⏳ Waiting for $URL to respond with \"$EXPECTED\"..."

for ((i=0; i<TIMEOUT; i++)); do
response=$(curl -s "$URL")
if [[ "$response" == "$EXPECTED" ]]; then
echo "✅ Service ready after $i seconds."
exit 0
fi
sleep 1
done

echo "❌ Timeout after $TIMEOUT seconds. Response: \"$response\""
exit 1

