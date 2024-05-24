import React, { FC } from 'react'
import Button from '@mui/material/Button'
import IconButton from '@mui/material/IconButton'
import ConstructionIcon from '@mui/icons-material/Construction'
import LoginIcon from '@mui/icons-material/Login'

import Cell from '../widgets/Cell'
import Row from '../widgets/Row'
import SessionModeSwitch from './SessionModeSwitch'
import SessionModeDropdown from './SessionModeDropdown'
import ModelPicker from './ModelPicker'

import useIsBigScreen from '../../hooks/useIsBigScreen'
import useAccount from '../../hooks/useAccount'

import {
  ISessionMode,
  ISessionType,
  SESSION_MODE_INFERENCE,
  SESSION_TYPE_TEXT,
} from '../../types'

const CreateToolbar: FC<{
  mode: ISessionMode,
  type: ISessionType,
  model: string,
  onOpenConfig: () => void,
  onSetMode: (mode: ISessionMode) => void,
  onSetModel: (model: string) => void,
}> = ({
  mode,
  type,
  model,
  onOpenConfig,
  onSetMode,
  onSetModel,
}) => {
  const bigScreen = useIsBigScreen()
  const account = useAccount()
  return (
    <Row>
      <Cell>
        {
          mode == SESSION_MODE_INFERENCE && type == SESSION_TYPE_TEXT && (
            <ModelPicker
              model={ model }
              onSetModel={ onSetModel }
            />
          )
        }
      </Cell>
      <Cell grow>
        
      </Cell>
      <Cell>
        <IconButton
          onClick={ onOpenConfig }
        >
          <ConstructionIcon />
        </IconButton>
      </Cell>
      <Cell>
        {
          bigScreen ? (
            <SessionModeSwitch
              mode={ mode }
              onSetMode={ onSetMode }
            />
          ) : (
            <SessionModeDropdown
              mode={ mode }
              onSetMode={ onSetMode }
            />
          )
        }
      </Cell>
      <Cell>
        {
          !account.user && (bigScreen ? (
            <Button
              size="medium"
              variant="contained"
              color="primary"
              endIcon={ <LoginIcon /> }
              onClick={ account.onLogin }
              sx={{
                ml: 2,
              }}
            >
              Login / Register
            </Button> 
          ) : (
            <IconButton
              onClick={ () => account.onLogin() }
            >
              <LoginIcon />
            </IconButton>
          ))
        }
      </Cell>
    </Row>
  )
}

export default CreateToolbar
