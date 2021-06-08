import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

type Data = {
  name: string;
};

const apiEndpoint = process.env.API_ENDPOINT;

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    const _res = await axios.get(`${apiEndpoint}/service`);
    res.status(200).json(_res.data);
  }

  if (req.method === 'POST') {
    const _res = await axios.post(`${apiEndpoint}/service`, req.body);
    res.status(200).json(_res.data);
  }

  if (req.method === 'PATCH') {
    const _res = await axios.patch(`${apiEndpoint}/service`, req.body);
    res.status(200).json(_res.data);
  }

  if (req.method === 'DELETE') {
    const _res = await axios.delete(`${apiEndpoint}/service`, {data: req.body});
    res.status(200).json(_res.data);
  }
};

export default handler;
