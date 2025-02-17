import mongoose from "mongoose";
import dotenv from "dotenv";

dotenv.config(); // Load .env again, just in case

const connectDB = async () => {
  try {
    if (!process.env.MONGODB_URI) {
      throw new Error("❌ MONGO_URI is not defined! Check your .env file.");
    }

    await mongoose.connect(process.env.MONGODB_URI);
    console.log("✅ MongoDB Connected Successfully!");
  } catch (error) {
    console.error("❌ MongoDB Connection Error:", error.message);
    process.exit(1);
  }
};

export default connectDB;

