import 'dotenv';
import { MongoClient } from 'mongodb';
import { CONNECTIONURL } from '$env/static/private';

if (!CONNECTIONURL) {
  throw new Error('CONNECTIONURL is not defined in the environment variables');
}

const client = new MongoClient(CONNECTIONURL);
const clientPromise = client.connect();

export default clientPromise;