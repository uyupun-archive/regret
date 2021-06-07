import type { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

type Data = {
  name: string;
};

const apiEndpoint = process.env.API_ENDPOINT;

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    await axios.get(`${apiEndpoint}/service`).then(_res => {
      res.status(200).json(_res.data);
    });
  }

  if (req.method === 'POST') {
    await axios.post(`${apiEndpoint}/service`, req.body).then(_res => {
      res.status(200).json(_res.data);
    });
  }

  if (req.method === 'PATCH') {
    await axios.patch(`${apiEndpoint}/service`, req.body).then(_res => {
      res.status(200).json(_res.data);
    });
  }

  if (req.method === 'DELETE') {
    await axios.delete(`${apiEndpoint}/service`, {data: req.body}).then(_res => {
      res.status(200).json(_res.data);
    });
  }
};

export default handler;
