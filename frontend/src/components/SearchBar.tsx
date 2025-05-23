import React from 'react'
import "../styles/SearchBar.css";

interface Props {
  onSearch: (searchTerm: string) => void;
}

const SearchBar: React.FC<Props> = ({ onSearch }: { onSearch: (searchTerm: string) => void }) => {

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    const element = document.getElementsByClassName("searchTerm")[0] as HTMLInputElement;
    element.value = value;
    onSearch(e.target.value);
  }

  const clearInput = () => {
    const element = document.getElementsByClassName("searchTerm")[0] as HTMLInputElement;
    element.value = "";
    onSearch("");
  }

  return (
    <div className='search-bar'>
      <input className="searchTerm" type="text" placeholder="Search character name..." onChange={handleInputChange} />
      <button className="clearBtn" type="button" onClick={clearInput}>Clear</button>
    </div>
  )
}

export default SearchBar