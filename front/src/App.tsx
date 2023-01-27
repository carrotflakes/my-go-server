import React, { useRef } from 'react';
import logo from './logo.svg';
import './App.css';

import { graphql } from '../src/gql';
import { useMutation, useQuery } from '@apollo/client';

const signinMutDocument = graphql(`
  mutation signin($email: String!, $password: String!) {
    signin(email: $email, password: $password)
  }
`)

const createNoteMutDocument = graphql(`
  mutation createNote($text: String!) {
    createNote(input: {text: $text}) {
      text
    }
  }
`)

const notesQueryDocument = graphql(`
  query notes {
    notes {
      text
    }
  }
`)

function App() {
  const emailEl = useRef(null as null | HTMLInputElement)
  const passwordEl = useRef(null as null | HTMLInputElement)
  const textEl = useRef(null as null | HTMLInputElement)
  const { data } = useQuery(notesQueryDocument, {})
  const [signin, signinRes] = useMutation(signinMutDocument)
  const [createNote, createNoteRes] = useMutation(createNoteMutDocument)

  return (
    <div className="App">
      <input type="text" placeholder='email' ref={emailEl} />
      <input type="text" placeholder='password' ref={passwordEl} />
      <button onClick={() => signin({ variables: { email: emailEl.current?.value || '', password: passwordEl.current?.value || '' } })}>signin</button>
      {JSON.stringify(signinRes.data)}

      <div>
      <input type="text" placeholder='text' ref={textEl} />
      <button onClick={() => createNote({ variables: { text: textEl.current?.value || '' } })}>add</button>
      </div>

      {JSON.stringify(data?.notes)}
    </div>
  );
}

export default App;
