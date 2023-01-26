import React from 'react';
import logo from './logo.svg';
import './App.css';

import { graphql } from '../src/gql';
import { useQuery } from '@apollo/client';

const notesQueryDocument = graphql(/* GraphQL */ `
  query notes {
    notes {
      text
    }
  }
`)

function App() {
  const { data } = useQuery(notesQueryDocument, {})

  return (
    <div className="App">
      {JSON.stringify(data?.notes)}
    </div>
  );
}

export default App;
