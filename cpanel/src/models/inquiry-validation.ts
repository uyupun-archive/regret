export interface InquiryValidation {
  id: number;
  service_id: number;
  is_required_subject: boolean;
  is_required_email: boolean;
  is_required_category: boolean;
  is_required_text: boolean;
};

export const initInquiryValidation = () => {
  return {
    id: 0,
    service_id: 0,
    is_required_subject: true,
    is_required_email: true,
    is_required_category: true,
    is_required_text: true
  };
};
