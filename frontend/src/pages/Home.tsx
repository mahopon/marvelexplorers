import React,{useState, useEffect, useRef} from "react";
import CharacterCard from "../components/CharacterCard.tsx";
import axios from "axios";
import "../styles/CharacterCard.css"

const DATA_API = "https://tcyao.duckdns.org/api/characters?offset=";

const Home = () => {
    const [characters, setCharacters] = useState([]);
    const [offset, setOffset] = useState(0);
    const allLoaded = useRef(false);
    const [loading, setLoading] = useState(true);

    const loadMore = (e) => {
        setOffset(offset + 20)
    }

    useEffect(() => {
        if (allLoaded.current) return;
        setLoading(true);
        axios.get(DATA_API+offset)
        .then((res) => {
            console.log(res.data);
            if (res.data.length === 0)
                allLoaded.current = true;
            setTimeout(() => {
                setCharacters([...characters,...res.data]);
                setLoading(false);
            }, 1000);
        }).catch((err) => {
            console.log(err);
            setLoading(false);
        });
    },[offset]);

    return (
        <>
            <header>
                <h1>Welcome to the Marvel Universe!</h1>
                <p>Explore marvel characters and their comics</p>
            </header>
            <section id="charSection">
                <h2>Choose your character!</h2>
                <div className="charTab">
                    { characters.map((char) => (
                        <CharacterCard
                          key={char.id}
                          name={char.name}
                          description={char.description}
                          resourceURI={char.resourceuri}
                          thumbnailPath={char.thumbnail_path}
                          thumbnailExtension={char.thumbnail_extension}
                        />
                    )) }
                    { /* Skeleton loading */
                        loading && Array(5).fill(0).map((_, index) => (
                            <div key={`skeleton-${index}`} className="character-card skeleton">
                                <div className="skeleton-img"></div>
                                <div className="skeleton-text"></div>
                            </div>
                        ))}
                </div>

                {!allLoaded.current && (
                    <button onClick={loadMore} disabled={loading}>
                        {loading ? "Loading..." : "Load More"}
                    </button>
                )}
            </section>
        </>
    )
}

export default Home;