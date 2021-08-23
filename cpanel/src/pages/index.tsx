import {useState, useEffect, Fragment} from 'react';
import {Modal} from 'react-bootstrap';
import axios from 'axios';
import {Service, AddService, initService, initAddService} from '../models/service';
import {Category, AddCategory, initCategory, initAddCategory} from '../models/category';

const Index = () => {
  const apiEndpoint = process.env.NEXT_PUBLIC_API_ENDPOINT;

  const [services, setServices] = useState<Array<Service>>([]);
  const [addingService, setAddingService] = useState<AddService>(initAddService());
  const [editingService, setEditingService] = useState<Service>(initService());
  const [deletingService, setDeletingService] = useState<Service>(initService());
  const [openedService, setOpenedService] = useState<Service | null>(null);
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  const [categories, setCategories] = useState<Array<Category>>([]);
  const [addingCategory, setAddingCategory] = useState<AddCategory>(initAddCategory());
  const [editingCategory, setEditingCategory] = useState<Category>(initCategory());

  const handleCloseDeleteServiceModal = () => setShowDeleteModal(false);
  const handleShowDeleteServiceModal = () => setShowDeleteModal(true);

  const fetchServices = () => {
    axios.get(`${apiEndpoint}/service`).then(res => {
      setServices(res.data);
    });
  };

  const addService = async (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    return await axios.post(`${apiEndpoint}/service`, addingService);
  };

  const editService = (service: Service) => {
    setEditingService(service);
  };

  const deleteService = async (service: Service) => {
    if (service.id === openedService?.id) setOpenedService(null);
    return await axios.delete(`${apiEndpoint}/service`, {data: service});
  };

  const openService = (service: Service) => {
    setOpenedService(service);
    setAddingCategory({...addingCategory, service_id: service.id});
    fetchCategories(service.id);
  };

  const saveService = async () => {
    setEditingService(initService());
    return await axios.patch(`${apiEndpoint}/service`, editingService);
  };

  const fetchCategories = (serviceId: number) => {
    axios.get(`${apiEndpoint}/category`, {
      params: {
        service_id: serviceId
      }
    }).then(res => {
      setCategories(res.data);
    });
  };

  const addCategory = (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    return axios.post(`${apiEndpoint}/category`, addingCategory);
  };

  const editCategory = (category: Category) => {
    setEditingCategory(category);
  };

  const deleteCategory = async (category: Category) => {
    return await axios.delete(`${apiEndpoint}/category`, {data: category});
  };

  const saveCategory = async () => {
    setEditingCategory(initCategory());
    return await axios.patch(`${apiEndpoint}/category`, editingCategory);
  };

  useEffect(fetchServices, []);

  return (
    <div className="container">
      <div className="mt-4 mb-4 d-flex align-items-center">
        <img src="logo.png" width="60" />
        <h1 className="ms-2">{ process.env.NEXT_PUBLIC_APP_NAME } コントロールパネル</h1>
      </div>
      <div className="mb-4">
        <h3>登録サービス一覧</h3>
        <form className="row g-1 mb-2" onSubmit={(e) => {
          addService(e).then(() => {
            setAddingService(initAddService());
            fetchServices();
          });
        }}>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名" value={addingService.name} onChange={(e) => {
              setAddingService({...addingService, name: e.target.value});
            }} />
          </div>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名（日本語）" value={addingService.name_ja} onChange={(e) => {
              setAddingService({...addingService, name_ja: e.target.value});
            }} />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="備考" value={addingService.description} onChange={(e) => {
              setAddingService({...addingService, description: e.target.value});
            }} />
          </div>
          <div className="col-2">
            <button className="btn btn-outline-primary">追加</button>
          </div>
        </form>
        {services.length <= 0 && (
          <p className="text-center h4 mt-4">サービスが登録されていません。</p>
        )}
        {services.length > 0 && (
          <table className="table">
            <thead>
              <tr>
                <th>ID</th>
                <th>サービス名</th>
                <th>サービス名（日本語）</th>
                <th>備考</th>
                <th>アクセストークン</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {services.map(service => {
                return (
                  <Fragment key={service.id}>
                    {editingService.id !== service.id &&
                      <tr>
                        <td>{service.id}</td>
                        <td>{service.name}</td>
                        <td>{service.name_ja}</td>
                        <td>{service.description}</td>
                        <td>{service.access_token}</td>
                        <td>
                          <div className="d-flex justify-content-around">
                            <button type="button" className="btn btn-outline-success btn-sm" onClick={() => editService(service)}>編集</button>
                            <button type="button" className="btn btn-outline-danger btn-sm" onClick={() => {
                              setDeletingService(service);
                              handleShowDeleteServiceModal();
                            }}>削除</button>
                            <button type="button" className="btn btn-outline-dark btn-sm" onClick={() => openService(service)}>開く</button>

                            <Modal
                              show={showDeleteModal}
                              onHide={handleCloseDeleteServiceModal}
                            >
                              <div className="modal-header">
                                <h5 className="modal-title">本当に削除しますか？</h5>
                              </div>
                              <div className="modal-body">一度削除すると後から戻すことはできません。</div>
                              <div className="modal-footer">
                                <button className="btn btn-outline-secondary" onClick={handleCloseDeleteServiceModal}>いいえ、削除しません</button>
                                <button className="btn btn-outline-danger" onClick={() => {
                                  deleteService(deletingService).then(() => {
                                    fetchServices();
                                    handleCloseDeleteServiceModal();
                                  });
                                }}>はい、削除します</button>
                              </div>
                            </Modal>
                          </div>
                        </td>
                      </tr>
                    }
                    {editingService.id === service.id &&
                      <tr>
                        <td>{editingService.id}</td>
                        <td>
                          <input type="text" className="form-control" value={editingService.name} onChange={(e) => {
                            setEditingService({...editingService, name: e.target.value});
                          }} />
                        </td>
                        <td>
                          <input type="text" className="form-control" value={editingService.name_ja} onChange={(e) => {
                            setEditingService({...editingService, name_ja: e.target.value});
                          }} />
                        </td>
                        <td>
                          <input type="text" className="form-control" value={editingService.description} onChange={(e) => {
                            setEditingService({...editingService, description: e.target.value});
                          }} />
                        </td>
                        <td>{editingService.access_token}</td>
                        <td>
                          <button type="button" className="btn btn-outline-success btn-sm" onClick={() => {
                            saveService().then(() => {
                              fetchServices();
                            });
                          }}>保存</button>
                        </td>
                      </tr>
                    }
                  </Fragment>
                )
              })}
            </tbody>
          </table>
        )}
      </div>
      {openedService && (
        <>
          <h3>▶ {openedService.id}: {openedService.name}({openedService.name_ja})</h3>
          <div className="mb-4">
            <h4>必須/任意項目切り替え</h4>
            <table className="table">
              <thead>
                <tr>
                  <th>項目名</th>
                  <th>必須/任意</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>件名</td>
                  <td>
                    <select className="form-select">
                      <option>必須</option>
                      <option>任意</option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <td>メールアドレス</td>
                  <td>
                    <select className="form-select">
                      <option>必須</option>
                      <option>任意</option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <td>カテゴリ</td>
                  <td>
                    <select className="form-select">
                      <option>必須</option>
                      <option>任意</option>
                    </select>
                  </td>
                </tr>
                <tr>
                  <td>本文</td>
                  <td>
                    <select className="form-select">
                      <option>必須</option>
                      <option>任意</option>
                    </select>
                  </td>
                </tr>
              </tbody>
            </table>
            <button type="button" className="btn btn-outline-success">保存</button>
          </div>
          <div>
            <h4>問い合わせカテゴリ一覧</h4>
            <form className="row g-1 mb-2" onSubmit={(e) => {
              addCategory(e).then(() => {
                const serviceId = openedService.id;
                fetchCategories(serviceId);
                setAddingCategory({
                  service_id: openedService.id,
                  name: '',
                  name_ja: ''
                });
              });
            }}>
              <div className="col-3">
                <input type="text" className="form-control" placeholder="カテゴリ名" value={addingCategory.name} onChange={(e) => {
                  setAddingCategory({...addingCategory, name: e.target.value});
                }} />
              </div>
              <div className="col-3">
                <input type="text" className="form-control" placeholder="カテゴリ名（日本語）" value={addingCategory.name_ja} onChange={(e) => {
                  setAddingCategory({...addingCategory, name_ja: e.target.value});
                }} />
              </div>
              <div className="col-2">
                <button className="btn btn-outline-primary">追加</button>
              </div>
            </form>
            {categories.length <= 0 && (
              <p className="text-center h4 mt-4">カテゴリが登録されていません。</p>
            )}
            {categories.length > 0 && (
              <table className="table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>カテゴリ名</th>
                    <th>カテゴリ名（日本語）</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  {categories.map(category => {
                    return (
                      <Fragment key={category.id}>
                        {editingCategory.id !== category.id &&
                            <tr>
                            <td>{category.id}</td>
                            <td>{category.name}</td>
                            <td>{category.name_ja}</td>
                            <td>
                              <button type="button" className="btn btn-outline-success btn-sm me-1" onClick={() => editCategory(category)}>編集</button>
                              <button type="button" className="btn btn-outline-danger btn-sm" onClick={() => {
                                deleteCategory(category).then(() => {
                                  fetchCategories(openedService.id);
                                });
                              }}>削除</button>
                            </td>
                          </tr>
                        }
                        {editingCategory.id === category.id &&
                          <tr>
                            <td>{editingCategory.id}</td>
                            <td>
                              <input type="text" className="form-control" value={editingCategory.name} onChange={(e) => {
                                setEditingCategory({...editingCategory, name: e.target.value});
                              }} />
                            </td>
                            <td>
                              <input type="text" className="form-control" value={editingCategory.name_ja} onChange={(e) => {
                                setEditingCategory({...editingCategory, name_ja: e.target.value});
                              }} />
                            </td>
                            <td>
                              <button type="button" className="btn btn-outline-success btn-sm" onClick={() => {
                                saveCategory().then(() => {
                                  fetchCategories(openedService.id);
                                });
                              }}>保存</button>
                            </td>
                          </tr>
                        }
                      </Fragment>
                    )
                  })}
                </tbody>
              </table>
            )}
          </div>
        </>
      )}
    </div>
  )
};

export default Index;
