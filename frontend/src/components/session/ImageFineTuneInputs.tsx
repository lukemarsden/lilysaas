import React, { FC, useState, useCallback, useEffect } from 'react'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Grid from '@mui/material/Grid'

import CloudUploadIcon from '@mui/icons-material/CloudUpload'
import ArrowCircleRightIcon from '@mui/icons-material/ArrowCircleRight'

import FileUpload from '../widgets/FileUpload'

import useAccount from '../../hooks/useAccount'
import Interaction from './Interaction'

import {
  SESSION_MODE_INFERENCE,
  SESSION_TYPE_TEXT,
  buttonStates,
} from '../../types'

import {
  getSystemMessage,
} from '../../utils/session'

export const ImageFineTuneInputs: FC<{
  initialFiles?: File[],
  onChange?: {
    (files: File[]): void
  },
  onDone: {
    (): void
  },
}> = ({
  initialFiles,
  onChange,
  onDone,
}) => {
  const account = useAccount()

  const [files, setFiles] = useState<File[]>(initialFiles || [])

  const onDropFiles = useCallback(async (newFiles: File[]) => {
    const existingFiles = files.reduce<Record<string, string>>((all, file) => {
      all[file.name] = file.name
      return all
    }, {})
    const filteredNewFiles = newFiles.filter(f => !existingFiles[f.name])
    setFiles(files.concat(filteredNewFiles))
  }, [
    files,
  ])

  useEffect(() => {
    if(!onChange) return
    onChange(files)
  }, [
    files,
  ])

  return (
    <Box
      sx={{
        mt: 2,
      }}
    >
      <Box
        sx={{
          mt: 4,
          mb: 4,
        }}
      >
        <Interaction
          session_id=""
          session_name=""
          interaction={ getSystemMessage('Firstly upload some images you want your model to learn from:') }
          type={ SESSION_TYPE_TEXT }
          mode={ SESSION_MODE_INFERENCE }
          serverConfig={ account.serverConfig }
        />
      </Box>
      <FileUpload
        sx={{
          width: '100%',
          mt: 2,
        }}
        onlyImages
        onUpload={ onDropFiles }
      >
        <Button
          sx={{
            width: '100%',
          }}
          variant="contained"
          color={ buttonStates.uploadFilesColor }
          endIcon={<CloudUploadIcon />}
        >
          { buttonStates.uploadFilesLabel }
        </Button>
        <Box
          sx={{
            border: '1px dashed #ccc',
            p: 2,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'flex-start',
            minHeight: '100px',
            cursor: 'pointer',
          }}
        >
          {
            files.length <= 0 && (
              <Typography
                sx={{
                  color: '#999',
                  width: '100%',
                }}
                variant="caption"
              >
                click or drop files here to upload them ...
              </Typography>
            )
          }
          <Grid container spacing={3} direction="row" justifyContent="flex-start">
            {
              files.length > 0 && files.map((file) => {
                const objectURL = URL.createObjectURL(file)
                return (
                  <Grid item xs={4} md={4} key={file.name}>
                    <Box
                      sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                        justifyContent: 'center',
                        color: '#999'
                      }}
                    >
                      <Box
                        component="img"
                        src={objectURL}
                        alt={file.name}
                        sx={{
                          height: '50px',
                          border: '1px solid #000000',
                          filter: 'drop-shadow(3px 3px 5px rgba(0, 0, 0, 0.2))',
                          mb: 1,
                        }}
                      />
                      <Typography variant="caption">
                        {file.name}
                      </Typography>
                      <Typography variant="caption">
                        ({file.size} bytes)
                      </Typography>
                    </Box>
                  </Grid>
                )
              })
                
            }
          </Grid>
        </Box>
      </FileUpload>
      {
        files.length > 0 && (
          <Button
            sx={{
              width: '100%',
            }}
            variant="contained"
            color="secondary"
            endIcon={<ArrowCircleRightIcon />}
            onClick={ onDone }
          >
            Next Step
          </Button>
        )
      }
    </Box>
  )   
}

export default ImageFineTuneInputs