import React, { useEffect, useState } from "react";
import { Link, useLocation } from "react-router-dom";

const SaveTicket = () => {
    const [id, setId] = useState(0);
    const [ticket, setTicket] = useState({
        title: "",
        description: "",
        contact: "",
        status: 0
    });

    //setup data for update
    const location = useLocation();
    useEffect(() => {
        if (location.state) {
            switch(location.state.status) {
                case "pending":
                    location.state.status = 0;
                    break;
                case "accepted":
                    location.state.status = 1;
                    break;
                case "accepted":
                    location.state.status = 2;
                    break;
                case "accepted":
                    location.state.status = 3;
                    break;
            }
            setId(location.state.id);
            setTicket({
                ...ticket,
                title: location.state.title,
                description: location.state.description,
                contact: location.state.contact,
                status: location.state.status
            });
        }
    },[])

    const handleChange = () => (e) => {
        let value = e.target.value;
        let name = e.target.name;
        setTicket({
            ...ticket,
            [name]: value,
        });
    }

    const onSubmit = e => {
        e.preventDefault();

        //const data = new FormData(e.target);
        //const payload = Object.fromEntries(data.entries());
        const myHeaders = new Headers();
        myHeaders.append("Content-Type","application/json");

        //update or create
        const requestOptions = {
            method: id==0?"POST":"PUT",
            body: JSON.stringify({
                "title": ticket.title,
                "description": ticket.description,
                "status": parseInt(ticket.status),
                "contact": ticket.contact,
            }),
            headers: myHeaders,
        };

        
        let fetchUrl = "";
        if(id != 0) {
            fetchUrl = "http://localhost:8000/tickets" + "/" + id;
        } else {
            fetchUrl = "http://localhost:8000/tickets"
        }

        fetch(fetchUrl, requestOptions)
            .then((response)=> response.json())
            .then(data => {
                console.log(data);
            });
        

        setTicket({
            title: "",
            description: "",
            contact: "",
            status: 0
        });
    }

    return <div className="w-screen m-3">
        <div className="flex">
            <div className="flex-1">
                <h1 className="p-2 text-xl capitalize">Add/Update Ticket</h1>
            </div>
        </div>
        <div className="p-8 mt-6 lg:mt-0 rounded shadow bg-white w-1/2">
            <form onSubmit={onSubmit}>
                <div className="mb-4">
                    <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="title">
                        Title
                    </label>
                    <input 
                        className="shadow appearance-none border focus:border-rose-400 rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
                        type="text" 
                        placeholder="title"
                        name="title" 
                        value={ticket.title}
                        onChange={handleChange("title")}
                        required
                    />
                </div>
                <div className="mb-6">
                    <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="description">
                        Description
                    </label>
                    <input 
                        className="shadow appearance-none border focus:border-rose-400 rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
                        type="text"
                        placeholder="description"
                        name="description" 
                        value={ticket.description}
                        onChange={handleChange("description")}
                        required
                    />
                </div>
                <div className="mb-6">
                    <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="contact">
                        Contact Infomation
                    </label>
                    <input 
                        className="shadow appearance-none border focus:border-rose-400 rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
                        type="text"
                        placeholder="contact info"
                        name="contact" 
                        value={ticket.contact}
                        onChange={handleChange("contact")}
                        required
                    />
                </div>
                <div className="inline-block relative w-full">
                    <select className="block appearance-none w-full bg-white border border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none focus:shadow-outline" 
                        name="status"
                        value={ticket.status}
                        onChange={handleChange("status")}
                    >
                        <option value={0}>pending</option>
                        <option value={1}>accepted</option>
                        <option value={2}>resolved</option>
                        <option value={3}>rejected</option>
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <svg className="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"/></svg>
                    </div>
                </div>
                <div className="flex items-center mt-5">
                    <button className="bg-rose-500 hover:bg-rose-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
                        Save
                    </button>
                    <Link to="/" className="ml-5 text-rose-400 hover:text-rose-600">
                        Cancel
                    </Link>
                </div>
            </form>
        </div>
    </div>;
}

export default SaveTicket