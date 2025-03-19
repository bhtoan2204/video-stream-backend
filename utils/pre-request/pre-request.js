// Secret key (must match the server's secret key)
const SECRET_KEY =
  "2b1080b6192d8c0fad991c54a6a3ae3ec60cf886ae4e9d4721d73d0e37713823";

// Function to generate HMAC SHA256 signature
function generateHMACSignature(message, secret) {
  const crypto = require("crypto-js");
  return crypto.HmacSHA256(message, secret).toString(crypto.enc.Hex);
}

// Generate timestamp (ISO 8601 format)
const timestamp = new Date().toISOString();

// Generate unique nonce
const nonce = `${Date.now()}-${Math.floor(Math.random() * 100000)}`;

// Compute HMAC signature using timestamp + nonce
const message = timestamp + nonce;
const signature = generateHMACSignature(message, SECRET_KEY);

// Set headers dynamically in Postman
pm.request.headers.add({ key: "X-Timestamp", value: timestamp });
pm.request.headers.add({ key: "X-Nonce", value: nonce });
pm.request.headers.add({ key: "X-Signature", value: signature });

// Log values for debugging (Optional)
console.log("X-Timestamp:", timestamp);
console.log("X-Nonce:", nonce);
console.log("X-Signature:", signature);
