import express from 'express';
const apiRouter = express.Router();

// Examples
//
// To get all movie details
// GET : localhost:PORT/movies
//
// To get specific movie details
// GET : localhost:PORT/movies?name=avatar

apiRouter.get('/movies',(req,res)=>{
    const { movieName }  = req.query;

    if(movieName) {
        console.log(movieName);
    }

    res.json({
        "movieName": "The Great Gatsby",
    })
})



export default apiRouter;
