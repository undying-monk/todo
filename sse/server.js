const express = require('express');
const cors = require('cors'); // Optional: Use if client and server are on different ports
const app = express();
const PORT = 3000;

app.use(cors());

app.get('/events', (req, res) => {
    // 1. Set required headers for SSE
    res.setHeader('Content-Type', 'text/event-stream');
    res.setHeader('Cache-Control', 'no-cache');
    res.setHeader('Connection', 'keep-alive');
    res.flushHeaders(); // Establish the stream immediately

    console.log('Client connected');

    // 2. Function to send formatted SSE data
    const sendData = (data) => {
        // SSE format requires "data: <content>\n\n"
        res.write(`data: ${JSON.stringify(data)}\n\n`);
    };

    // 3. Send periodic updates
    const intervalId = setInterval(() => {
        const message = {
            time: new Date().toLocaleTimeString(),
            info: "Server update"
        };
        sendData(message);
    }, 2000);

    // 4. Clean up when the client closes the connection
    req.on('close', () => {
        console.log('Client disconnected');
        clearInterval(intervalId);
        res.end();
    });
});

app.listen(PORT, () => {
    console.log(`SSE Server running at http://localhost:${PORT}/events`);
});