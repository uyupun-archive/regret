import type { NextApiRequest, NextApiResponse } from 'next';

type Data = {
  msg: string;
};

const handler = async (req: NextApiRequest, res: NextApiResponse<Data>) => {
  if (req.method === 'GET') {
    return res.status(200).json({msg: 'pong'});
  }
};

export default handler;
