import { useEffect, useState } from "react";

function App() {
  const [data, setData] = useState(undefined);

  useEffect(() => {
    (async () => {
      const raw = await fetch("http://localhost:8080/api/info");
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
}

export default App;
