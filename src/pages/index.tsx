const Index = () => {
  return (
    <div className="container">
      <div className="mt-4 mb-4 d-flex align-items-center">
        <img src="logo.png" width="60" />
        <h1 className="ms-2">{ process.env.NEXT_PUBLIC_APP_NAME } コントロールパネル</h1>
      </div>
      <div className="mb-4">
        <h3>登録サービス一覧</h3>
        <form className="row g-1 mb-2">
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名" />
          </div>
          <div className="col-3">
            <input type="text" className="form-control" placeholder="サービス名（日本語）" />
          </div>
          <div className="col-4">
            <input type="text" className="form-control" placeholder="備考" />
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
            <tr>
              <td>1</td>
              <td>official</td>
              <td>オフィシャル</td>
              <td>オフィシャルサイトです</td>
              <td>xxxxxxxxxxxxxxxxx</td>
              <td>
                <button type="button" className="btn btn-outline-success btn-sm">編集</button>
                <button type="button" className="btn btn-outline-danger btn-sm ms-1">削除</button>
                <button type="button" className="btn btn-outline-dark btn-sm ms-1">開く</button>
              </td>
            </tr>
            <tr>
              <td>2</td>
              <td>takakeibo</td>
              <td>たかしのカケイボ</td>
              <td>家計簿アプリです</td>
              <td>xxxxxxxxxxxxxxxxx</td>
              <td>
                <button type="button" className="btn btn-outline-success btn-sm">編集</button>
                <button type="button" className="btn btn-outline-danger btn-sm ms-1">削除</button>
                <button type="button" className="btn btn-outline-dark btn-sm ms-1">開く</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div>
        <h3>問い合わせカテゴリ一覧</h3>
        <h4>▶ official</h4>
        <table className="table">
          <thead>
            <tr>
              <td>ID</td>
              <td>カテゴリ名</td>
              <td>カテゴリ名（日本語）</td>
              <td>操作</td>
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
