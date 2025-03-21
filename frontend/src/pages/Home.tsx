import React,{useState, useEffect, useRef} from "react";
import CharacterCard from "../components/CharacterCard.tsx";
import axios from "axios";

const DATA_API = "http://139.59.126.66/characters?offset=";

const Home = () => {
    const [characters, setCharacters] = useState([]);
    const offset = useRef(0);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        axios.get(DATA_API+offset.current)
        .then((res) => {
            console.log(res);
            // setCharacters(res)
        })
    },offset)
    return (
        <div>
            <h1>Welcome to the Marvel Universe!</h1>
            <p>Explore marvel characters and their comics</p>
            <CharacterCard
              name="3-D Man"
              description=""
              resourceURI="http://gateway.marvel.com/v1/public/characters/1011334"
              thumbnailPath="http://i.annihil.us/u/prod/marvel/i/mg/c/e0/535fecbbb9784"
              thumbnailExtension="jpg"
            />
        </div>
    )
}

export default Home;