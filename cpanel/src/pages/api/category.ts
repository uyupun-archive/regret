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
    res.status(200).json(_res.data);
  }

  if (req.method === 'POST') {
    const _res = await axios.post(`${apiEndpoint}/category`, req.body);
    res.status(200).json(_res.data);
  }

  if (req.method === 'PATCH') {
    const _res = await axios.patch(`${apiEndpoint}/category`, req.body);
    res.status(200).json(_res.data);
  }

  if (req.method === 'DELETE') {
    const _res = await axios.delete(`${apiEndpoint}/category`, {data: req.body});
    res.status(200).json(_res.data);
  }
};

export default handler;
