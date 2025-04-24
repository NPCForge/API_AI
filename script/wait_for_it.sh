URL="http://localhost:3000/health"
EXPECTED="OK"
TIMEOUT=30  # timeout en secondes

echo "⏳ Attente que $URL réponde avec \"$EXPECTED\"..."

for ((i=0; i<TIMEOUT; i++)); do
response=$(curl -s "$URL")
echo response
if [[ "$response" == "$EXPECTED" ]]; then
echo "✅ Service prêt après $i secondes."
exit 0
fi
sleep 1
done

echo "❌ Timeout après $TIMEOUT secondes. Réponse: \"$response\""
exit 1

