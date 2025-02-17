import mongoose from "mongoose";
const { schema } = mongoose;

// Define the movie schema
const movieSchema = new schema({
    title: {
        type: String,
        required: true,
        trim: true
    },
    year: {
        type: Number,
        required: true
    },
    rating: {
        type: Number,
        required: true
    },
    runTime: {
        type: String,
        required: true
    },
    pg: {
        type: String,
        required: true
    },
    genre: [{
        type: String,
        enum: ['Action', 'Comedy', 'Drama', 'Horror', 'Sci-Fi', 'Romance', 'Thriller', 'Documentary'],
        required: true
    }],
    description: {
        type: String,
        required: true
    },
    poster: {
        // url format
        type: String,
        required: true
    },
    background: {
        // url format
        type: String,
        required: true
    }
})

// text index for search functionality
movieSchema.index({ name: 'text', description: 'text' });

// to search movie by genre
movieSchema.statics.findByGenre = function(genre) {
    return this.find({ genre: genre });
};

const Movie = mongoose.model('Movie', movieSchema);
export default Movie;
