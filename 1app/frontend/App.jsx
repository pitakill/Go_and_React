const App = () => {
  const [data, setData] = React.useState(undefined);

  React.useEffect(() => {
    (async () => {
      const raw = await fetch("/api/info");
      const data = await raw.json();

      setData(data);
    })();
  }, []);

  if (data === undefined) {
    return null;
  }

  return (
    <main className="container">
      <article>
        <header>{data.id}</header>
        <p>Nombre: {data.name}</p>
        <p>Email: {data.email}</p>
        Lenguajes:
        <ul>
          {data.languages.map((e, i) => (
            <li key={i}>{e}</li>
          ))}
        </ul>
      </article>
    </main>
  );
};

const root = ReactDOM.render(<App />, document.getElementById("root"));
