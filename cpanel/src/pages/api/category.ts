import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

type Data = {
  name: string;
};

const apiEndpoint = process.env.API_ENDPOINT;

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    const _res = await axios.get(`${apiEndpoint}/category`, {
      params: {
        service_id: req.query.service_id
      }
    })
    return res.status(200).json(_res.data);
  }

  else if (req.method === 'POST') {
    const _res = await axios.post(`${apiEndpoint}/category`, req.body);
    return res.status(200).json(_res.data);
  }

  else if (req.method === 'PATCH') {
    const _res = await axios.patch(`${apiEndpoint}/category`, req.body);
    return res.status(200).json(_res.data);
  }

  else if (req.method === 'DELETE') {
    const _res = await axios.delete(`${apiEndpoint}/category`, {data: req.body});
    return res.status(200).json(_res.data);
  }

  else {
    res.setHeader('Allow', ['GET', 'POST', 'PATCH', 'DELETE']);
    return res.status(405).end(`Method ${req.method} Not Allowed`);
  }
};

export default handler;
