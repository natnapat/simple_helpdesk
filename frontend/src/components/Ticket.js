import React, { useEffect, useState, Fragment } from "react";
import { Menu, Transition } from '@headlessui/react'
import { IoIosRefresh } from "react-icons/io";
import { AiFillCaretDown } from "react-icons/ai"
import { FiEdit2 } from "react-icons/fi";
import { Link } from "react-router-dom";

const Ticket = () => {
    const [items, setItems] = useState([]);
    const [status, setStatus] = useState("");

    if (status != "") {
        useEffect(() => {
            fetch("http://localhost:8000/tickets/" + status)
                .then(res => res.json())
                .then(
                    (result) => {
                        //console.log(result);
                        setItems(result);
                    }
                )
        },[status])
    } else {
        useEffect(() => {
            fetch("http://localhost:8000/tickets")
                .then(res => res.json())
                .then(
                    (result) => {
                        //console.log(result);
                        setItems(result);
                    }
                )
        },[status])
    }

    function refreshPage() {
        window.location.reload(false);
    }
    
    return (
        <div className="w-screen m-3">
            <div className="flex">
                <div className="flex-1">
                    <h1 className="p-2 text-xl capitalize">{status==""?"All":status} Tickets ({items.length})</h1>
                </div>
                <div className="flex-end">
                    <Menu as="div" className="relative inline-block text-left">
                        <div>
                            <Menu.Button className="justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50">
                                Status
                                <AiFillCaretDown size={12} className="inline-flex cursor-pointer ml-1"/>
                            </Menu.Button>
                        </div>
                        <Transition
                        as={Fragment}
                        enter="transition ease-out duration-100"
                        enterFrom="transform opacity-0 scale-95"
                        enterTo="transform opacity-100 scale-100"
                        leave="transition ease-in duration-75"
                        leaveFrom="transform opacity-100 scale-100"
                        leaveTo="transform opacity-0 scale-95"
                        >
                            <Menu.Items className="absolute right-0 mt-2 w-24 origin-top-right divide-y divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
                                <div className="px-1 py-1 ">
                                    <Menu.Item>
                                        {({ active }) => (
                                        <button
                                            className={`${
                                            active ? 'bg-yellow-500 text-white' : 'text-gray-900'
                                            } group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                                            onClick={() => setStatus("pending")}
                                        >
                                            pending
                                        </button>
                                        )}
                                    </Menu.Item>
                                    <Menu.Item>
                                        {({ active }) => (
                                        <button
                                            className={`${
                                            active ? 'bg-sky-500 text-white' : 'text-gray-900'
                                            } group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                                            onClick={() => setStatus("accepted")}
                                        >
                                            accepted
                                        </button>
                                        )}
                                    </Menu.Item>
                                    <Menu.Item>
                                        {({ active }) => (
                                        <button
                                            className={`${
                                            active ? 'bg-green-500 text-white' : 'text-gray-900'
                                            } group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                                            onClick={() => setStatus("resolved")}
                                        >
                                            resolved
                                        </button>
                                        )}
                                    </Menu.Item>
                                    <Menu.Item>
                                        {({ active }) => (
                                        <button
                                            className={`${
                                            active ? 'bg-red-500 text-white' : 'text-gray-900'
                                            } group flex w-full items-center rounded-md px-2 py-2 text-sm`}
                                            onClick={() => setStatus("rejected")}
                                        >
                                            rejected
                                        </button>
                                        )}
                                    </Menu.Item>
                                </div>
                            </Menu.Items>
                        </Transition>
                    </Menu>
                </div>
                <div className="flex-end mx-3">
                    <button className="bg-white text-gray-600 hover:text-black hover:bg-gray-50 py-2 px-2 rounded border border-gray-300" onClick={refreshPage}>
                        <IoIosRefresh
                            size={20}
                            className="cursor-pointer"
                            onClick={() => setOpen(!open)}
                        />
                    </button>
                </div>   
                <div className="flex-none">
                    <Link to="/save">
                        <button className="flex-none bg-rose-400 hover:bg-rose-500 text-white font-bold py-2 px-2 rounded">
                            Add ticket
                        </button>
                    </Link>
                </div>
            </div>
            <div className="p-8 mt-6 lg:mt-0 rounded shadow bg-white overflow-auto">
                <table>
                    <thead className="bg-gray-50 border-b-2 border-gray-200">
                        <tr>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Title</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Description</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Contact Info</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Status</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Last Update</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Created At</th>
                            <th className="p-3 text-sm font-semibold tracking-wide text-left">Action</th>
                        </tr>
                    </thead>
                    <tbody>
                        {items.map((row) => ( 
                            <tr key={row.id} className="hover:bg-rose-50">
                                <td className="p-3 text-sm text-gray-700">{row.title}</td>
                                <td className="p-3 text-sm text-gray-700">{row.description}</td>
                                <td className="p-3 text-sm text-gray-700">{row.contact}</td>
                                <td className="p-3 text-sm text-gray-700">
                                    <span className={row.status}>{row.status}</span>
                                </td>
                                <td className="p-3 text-sm text-gray-700">{row.updated_at}</td>
                                <td className="p-3 text-sm text-gray-700">{row.created_at}</td>
                                <td className="p-3 text-sm text-gray-700">
                                    <Link to="/save" state={{id:row.id, title:row.title, description:row.description, contact:row.contact, status:row.status}}>
                                        <FiEdit2 size={20} className="inline-flex cursor-pointer ml-1 hover:text-rose-400"/>
                                    </Link>
                                </td>
                            </tr>
                            
                        ))} 
                    </tbody>
                </table>
            </div>
                
        </div>
    );
}

export default Ticket