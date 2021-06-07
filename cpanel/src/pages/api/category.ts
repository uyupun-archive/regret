import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

type Data = {
  name: string;
};

const apiEndpoint = process.env.API_ENDPOINT;

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    await axios.get(`${apiEndpoint}/category`, {
      params: {
        service_id: req.query.service_id
      }
    }).then(_res => {
      res.status(200).json(_res.data);
    });
  }
};

export default handler;
