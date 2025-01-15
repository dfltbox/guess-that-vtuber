import type { RequestHandler } from '@sveltejs/kit';
import clientPromise from '../../../../../lib/db';

export const GET: RequestHandler = async ({ params }) => {
  const client = await clientPromise;
  const collection = client.db("vtuberdb").collection("vtubers");
  const org = decodeURIComponent(params.org || '');
  const vtubers = await collection.find({ org: org }).toArray();

  return new Response(JSON.stringify(vtubers), {
    status: 200,
    headers: {
      'Content-Type': 'application/json'
    }
  });
};