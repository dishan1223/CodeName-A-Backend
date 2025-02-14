// Dependencies

import dotenv from 'dotenv';
dotenv.config();
import express from 'express';
import bodyParser from 'body-parser';
import cors from 'cors';

// Import Routes
import indexRouter from './routes/index.route.js';
import apiRouter from './routes/api.route.js';

const app = express();

// .env
const PORT = process.env.PORT || 3000;

// middleware
app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// set routes
app.use('/', indexRouter);
app.use('/api', apiRouter);


app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
