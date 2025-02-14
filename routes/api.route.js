import express from 'express';
const apiRouter = express.Router();

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
