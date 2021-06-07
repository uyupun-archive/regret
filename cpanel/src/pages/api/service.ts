import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

type Data = {
  name: string;
};

const apiEndpoint = process.env.API_ENDPOINT;

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    await axios.get(`${apiEndpoint}/service`).then(_res => {
      res.status(200).json(_res.data)
    })
  }

  if (req.method === 'POST') {
    await axios.post(`${apiEndpoint}/service`, req.body).then(() => {
      res.status(200).json({name: "test"})
    })
  }
};

export default handler;
