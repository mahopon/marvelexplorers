import React,{useState, useEffect, useRef} from "react";
import CharacterCard from "../components/CharacterCard.tsx";
import axios from "axios";
import "../styles/animations.css";
import "../styles/CharacterCard.css";
import {Character} from "../interfaces/CharacterInterface.tsx"
import CharacterDetails from "../components/CharacterDetails.tsx";

const DATA_API = "https://tcyao.duckdns.org/api/characters?offset=";

const Home = () => {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [offset, setOffset] = useState(0);
    const allLoaded = useRef(false);
    const [loading, setLoading] = useState(true);
    const [selectedChar, setSelectedChar] = useState<Character | undefined>(undefined)

    const loadMore = (e) => {
        setOffset(offset + 20)
    }

    const selectChar = (char: Character) => {
        setSelectedChar(char);
        console.log(char)
    }

    useEffect(() => {
        if (allLoaded.current) return;
        setLoading(true);
        axios.get(DATA_API+offset)
        .then((res) => {
            console.log(res.data);
            if (res.data.length === 0)
                allLoaded.current = true;
            let newChars: Character[] = [];
            setTimeout(() => {
                for (let newChar of res.data) {
                    newChars.push({
                        id: newChar.id,
                        name: newChar.name,
                        resourceURI: newChar.resourceuri,
                        thumbnailExtension: newChar.thumbnail_extension,
                        thumbnailPath: newChar.thumbnail_path
                    })
                }
                setCharacters([...characters,...newChars]);
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
                          character={char}
                          onClick={selectChar}
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
            {selectedChar && (
                <CharacterDetails
                  character = {selectedChar}
                  setSelectedChar={setSelectedChar}
                ></CharacterDetails>
            )}
        </>
    )
}

export default Home;