import gql from 'graphql-tag';
import * as React from 'react';
import { GetPxScripts, Script } from 'utils/script-bundle';

import { useApolloClient } from '@apollo/react-hooks';

const GET_USER_ORG = gql`
{
  user {
    orgName
  }
}
`;

interface ScriptsContextProps {
  scripts: Script[];
  promise: Promise<Script[]>;
}

export const ScriptsContext = React.createContext<ScriptsContextProps>(null);

export const ScriptsContextProvider = (props) => {
  const client = useApolloClient();
  const [scripts, setScripts] = React.useState<Script[]>([]);

  const promise = React.useMemo(() => {
    return client.query({ query: GET_USER_ORG, fetchPolicy: 'network-only' })
      .then((result) => {
        const orgName = result?.data?.user.orgName;
        return GetPxScripts(orgName);
      });
  }, []);

  React.useEffect(() => {
    // Do this only once.
    promise.then(setScripts);
  }, []);

  const context = React.useMemo(() => ({ scripts, promise }), [scripts]);

  return (
    <ScriptsContext.Provider value={context}>
      {props.children}
    </ScriptsContext.Provider>
  );
};