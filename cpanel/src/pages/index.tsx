import {useState, Fragment} from 'react';
import axios from 'axios';
import {Service, AddService, initAddService} from '../models/service';

const Index = () => {
  const [services, setServices] = useState<Array<Service>>([
    {
      id: 1,
      name: 'official',
      name_ja: 'オフィシャル',
      description: 'オフィシャルサイトです',
      access_token: 'xxxxxxxxxxxxxxxxx'
    },
    {
      id: 2,
      name: 'takakeibo',
      name_ja: 'たかしのカケイボ',
      description: '家計簿アプリです',
      access_token: 'xxxxxxxxxxxxxxxxx'
    }
  ]);

  const [addingService, setAddingService] = useState<AddService>(initAddService());
  const [editingService, setEditingService] = useState<Service | object>({});

  const addService = async (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    const res = await axios.post('http://localhost:3000/api/service', addingService);
  };

  const editService = (service: Service) => {
    setEditingService(service);
  };

  const deleteService = () => {

  };

  const openService = () => {

  };

  const saveService = () => {
    setEditingService({});
  };

  const addCategory = () => {

  };

  const editCategory = () => {

  };

  const deleteCategory = () => {

  };

  return (
    <div className="container">
      <div className="mt-4 mb-4 d-flex align-items-center">
        <img src="logo.png" width="60" />
        <h1 className="ms-2">{ process.env.NEXT_PUBLIC_APP_NAME } コントロールパネル</h1>
      </div>
      <div className="mb-4">
        <h3>登録サービス一覧</h3>
        <form className="row g-1 mb-2" onSubmit={addService}>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名" onChange={(e) => {
              setAddingService({...addingService, name: e.target.value});
            }} />
          </div>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名（日本語）" onChange={(e) => {
              setAddingService({...addingService, name_ja: e.target.value});
            }} />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="備考" onChange={(e) => {
              setAddingService({...addingService, description: e.target.value});
            }} />
          </div>
          <div className="col-2">
            <button className="btn btn-outline-primary">追加</button>
          </div>
        </form>
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
            {services.map((service, idx) => {
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
                        <button type="button" className="btn btn-outline-success btn-sm" onClick={() => {editService(service)}}>編集</button>
                        <button type="button" className="btn btn-outline-danger btn-sm ms-1">削除</button>
                        <button type="button" className="btn btn-outline-dark btn-sm ms-1">開く</button>
                      </td>
                    </tr>
                  }
                  {editingService.id === service.id &&
                    <tr>
                      <td>{service.id}</td>
                      <td>
                        <input type="text" className="form-control" defaultValue={service.name} />
                      </td>
                      <td>
                        <input type="text" className="form-control" defaultValue={service.name_ja} />
                      </td>
                      <td>
                        <input type="text" className="form-control" defaultValue={service.description} />
                      </td>
                      <td>{service.access_token}</td>
                      <td>
                        <button type="button" className="btn btn-outline-success btn-sm" onClick={saveService}>保存</button>
                      </td>
                    </tr>
                  }
                </Fragment>
              )
            })}
          </tbody>
        </table>
      </div>
      <div>
        <h3>問い合わせカテゴリ一覧</h3>
        <h4>▶ official</h4>
        <form className="row g-1 mb-2">
          <div className="col-3">
            <input type="text" className="form-control" placeholder="カテゴリ名" />
          </div>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="カテゴリ名（日本語）" />
          </div>
          <div className="col-2">
            <button className="btn btn-outline-primary">追加</button>
          </div>
        </form>
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
            <tr>
              <td>1</td>
              <td>bug</td>
              <td>バグ報告</td>
              <td>
                <button type="button" className="btn btn-outline-success btn-sm">編集</button>
                <button type="button" className="btn btn-outline-danger btn-sm ms-1">削除</button>
              </td>
            </tr>
            <tr>
              <td>2</td>
              <td>opinion</td>
              <td>意見</td>
              <td>
                <button type="button" className="btn btn-outline-success btn-sm">編集</button>
                <button type="button" className="btn btn-outline-danger btn-sm ms-1">削除</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
};

export default Index;
