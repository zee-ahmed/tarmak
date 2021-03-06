.. _tarmak_kubeconfig:

tarmak kubeconfig
-----------------

Verify and print path to Kubeconfig

Synopsis
~~~~~~~~


Verify and print path to Kubeconfig

::

  tarmak kubeconfig [flags]

Options
~~~~~~~

::

  -h, --help          help for kubeconfig
  -p, --path string   Path to store kubeconfig file (default "${TARMAK_CONFIG}/${CURRENT_CLUSTER}/kubeconfig")

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

  -c, --config-directory string        config directory for tarmak's configuration (default "~/.tarmak")
      --current-cluster string         override the current cluster set in the config
      --keep-containers                do not clean-up terraform/packer containers after running them
      --log-flush-frequency duration   Maximum number of seconds between log flushes (default 5s)
      --public-api-endpoint            Override kubeconfig to point to cluster's public API endpoint
  -v, --verbose                        enable verbose logging
      --wing-dev-mode                  use a bundled wing version rather than a tagged release from GitHub

SEE ALSO
~~~~~~~~

* `tarmak <tarmak.html>`_ 	 - Tarmak is a toolkit for provisioning and managing Kubernetes clusters.

