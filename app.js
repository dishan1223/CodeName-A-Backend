// Dependencies

import dotenv from 'dotenv';
dotenv.config();

import express from 'express';
import bodyParser from 'body-parser';
import cors from 'cors';
import connectDB from './db/db.js';

// Import Routes
import indexRouter from './routes/index.route.js';
import apiRouter from './routes/api.route.js';

// initiate app
const app = express();
const PORT = process.env.PORT || 3000;

// middleware
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

connectDB();


// set routes
app.use('/', indexRouter);
app.use('/api', apiRouter);


app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
