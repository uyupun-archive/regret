export interface Category {
  id: number;
  service_id: number;
  name: string;
  name_ja: string;
};

export const initCategory = () => {
  return {
    id: 0,
    service_id: 0,
    name: '',
    name_ja: ''
  };
};

export interface AddCategory {
  service_id: number;
  name: string;
  name_ja: string;
};

export const initAddCategory = () => {
  return {
    service_id: 0,
    name: '',
    name_ja: ''
  };
};
