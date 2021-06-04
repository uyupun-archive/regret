export interface Service {
  id: number;
  name: string;
  name_ja: string;
  description: string;
  access_token: string;
};

export interface AddService {
  name: string;
  name_ja: string;
  description: string;
}

export const initAddService = (): AddService => {
  return {
    name: '',
    name_ja: '',
    description: ''
  }
}
