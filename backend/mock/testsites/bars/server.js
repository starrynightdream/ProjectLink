// server.js
const express = require('express');
const app = express();
const path = require('path');
const PORT = 8888;

// 中间件
app.use(express.static('public'));
app.use(express.json());

// 初始数据
let chartData = Array.from({length: 20}, () => Math.floor(Math.random() * 100) + 1);

// 路由
app.get('/display', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/display.html'));
});

app.get('/control', (req, res) => {
    res.sendFile(path.join(__dirname, 'public/control.html'));
});

// API
app.get('/api/data', (req, res) => {
    res.json(chartData);
});

app.post('/api/update', (req, res) => {
    chartData.shift();
    chartData.push(Math.floor(Math.random() * 100) + 1);
    res.json(chartData);
});

app.listen(PORT, (err) => {
    console.log(`Server running at http://localhost:${PORT}`);
    if (err) console.log(err)
});